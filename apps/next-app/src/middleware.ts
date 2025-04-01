import { Hono } from 'hono';
import type { NextRequest } from 'next/server';
import { NextResponse } from 'next/server';

import type { SessionData } from '@auth0/nextjs-auth0/types';
import { auth0 } from './lib/auth0';

interface Variables {
	session: SessionData;
}

const app = new Hono<{ Variables: Variables }>();

// auth0のMiddleware処理
app.all('/auth/*', async (c) => {
	return await auth0.middleware(c.req.raw as NextRequest);
});

app.all('/', (_c) => {
	return NextResponse.next();
});
app.use('*', async (c, next) => {
	const session = await auth0.getSession(c.req.raw as NextRequest);
	if (!session) {
		return NextResponse.redirect('/login');
	}
	c.set('session', session);
	await next();
});

// Middlewareはこれより上に書く
app.all('*', (c) => {
	const req = c.req.raw as NextRequest;
	return NextResponse.next({
		request: req,
	});
});

export async function middleware(request: NextRequest) {
	return app.fetch(request);
}

export const config = {
	matcher: [
		/*
		 * Match all request paths except for the ones starting with:
		 * - _next/static (static files)
		 * - _next/image (image optimization files)
		 * - favicon.ico, sitemap.xml, robots.txt (metadata files)
		 */
		'/((?!_next/static|_next/image|favicon.ico|sitemap.xml|robots.txt).*)',
	],
};

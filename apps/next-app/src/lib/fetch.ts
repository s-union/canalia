import createClient, { type Middleware } from 'openapi-fetch';
import type { paths } from '../../generated/schema';
import { auth0 } from './auth0';

const client = createClient<paths>({
	baseUrl: process.env.NEXT_PUBLIC_API_BASE_URL,
});

const fetchMiddleware: Middleware = {
	async onRequest({ request }) {
		const token = await auth0.getAccessToken();
		request.headers.set('Authorization', `Bearer ${token.token}`);
		return request;
	},
};

client.use(fetchMiddleware);

export default client;

import { auth0 } from '@/lib/auth0';
import client from '@/lib/fetch';

export default async function Home() {
	const session = await auth0.getSession();

	if (session) {
		const { data } = await client.GET('/user');

		return (
			<main>
				<h1>Welcome, {session.user.name}!</h1>
				<div>Access Token: {session.tokenSet.accessToken}</div>
				<div>ID Token: {session.tokenSet.idToken}</div>
				<a href="/auth/logout">Log out</a>
				<h2>API Data</h2>
				<div>name: {data?.name}</div>
				<div>email: {data?.email}</div>
			</main>
		);
	}

	return (
		<main>
			<div>
				<a href="/auth/login?screen_hint=signup">Sign up</a>
			</div>
			<div>
				<a href="/auth/login">Log in</a>
			</div>
		</main>
	);
}

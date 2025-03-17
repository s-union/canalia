import { client } from '@/lib/fetch';

export default async function Home() {
	const { data, error } = await client.GET('/user');
	if (error) {
		return <h1>Error fetching user data</h1>;
	}
	return <h1>Hello {data.name}!</h1>;
}

import createClient from 'openapi-fetch';
import type { paths } from '../../generated/schema';

export const client = createClient<paths>({
	baseUrl: process.env.NEXT_PUBLIC_API_BASE_URL,
});

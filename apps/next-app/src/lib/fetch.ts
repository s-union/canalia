import createClient from 'openapi-fetch';
import type { paths } from '../../generated/schema';

export const client = createClient<paths>({
	baseUrl: process.env.API_BASE_URL,
});

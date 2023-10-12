import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	return {
		shop: params.shop,
		item: params.item
	};
};

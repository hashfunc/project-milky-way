interface Response<T, U> {
	code: string;
	message: string;
	meta?: U;
	data: T;
}

interface SearchMeta {
	total_count: number;
	pagable_count: number;
	is_end: boolean;
}

export interface SearchData {
	id: string;
	place_name: string;
	category_name: string;
	category_group_code: string;
	category_group_name: string;
	phone: string;
	address_name: string;
	road_address_name: string;
	x: string;
	y: string;
	place_url: string;
	distance: string;
}

export type SearchResponse = Response<SearchData[], SearchMeta>;

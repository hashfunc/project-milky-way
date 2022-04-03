interface Star {
	name: string;
	address_name: string;
	road_address_name: string;
	latitude: string;
	longitude: string;
}

export type StarResponse = DefaultResponse<Star[]>;

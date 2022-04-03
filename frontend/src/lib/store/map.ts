import { writable } from 'svelte/store';

interface MapStore {
	map?: kakao.maps.Map;
	latitude: string;
	longitude: string;
	markers: kakao.maps.Marker[];
}

const initial: MapStore = {
	latitude: '33.450701',
	longitude: '126.570667',
	markers: []
};

export const store = writable(initial);

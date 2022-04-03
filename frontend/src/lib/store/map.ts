import { writable } from 'svelte/store';

interface MapStore {
	latitude: string;
	longitude: string;
}

const initial: MapStore = {
	latitude: '33.450701',
	longitude: '126.570667'
};

export const store = writable(initial);

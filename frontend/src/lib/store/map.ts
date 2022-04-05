import { get, writable } from 'svelte/store';
import IconStar from '$lib/assets/icon-star.png';

interface MapStore {
	map?: kakao.maps.Map;
	marker?: kakao.maps.Marker;
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

interface Position {
	latitude: string;
	longitude: string;
}

export const actions = {
	initialize(ref: HTMLElement) {
		const { latitude, longitude } = get(store);

		const position = new kakao.maps.LatLng(latitude, longitude);

		const options = { center: position, level: 4 };
		const map = new kakao.maps.Map(ref, options);

		const marker = new kakao.maps.Marker({ map, position });

		store.update((prev) => ({ ...prev, map, marker }));
	},
	setCenter(location: Position) {
		store.update((prev) => ({ ...prev, ...location }));

		const { map, marker } = get(store);

		const center = new kakao.maps.LatLng(location.latitude, location.longitude);

		map?.setCenter(center);
		marker?.setPosition(center);
	},
	updateMarkers(locations: Position[]) {
		const { latitude, longitude, map, markers: prevMarkers } = get(store);

		prevMarkers.forEach((marker) => {
			marker.setMap(null);
		});

		const markers = locations.map((location) => {
			const image = new kakao.maps.MarkerImage(IconStar, new kakao.maps.Size(32, 32));
			const position = new kakao.maps.LatLng(location.latitude, location.longitude);
			return new kakao.maps.Marker({ map, position, image });
		});

		const center = new kakao.maps.LatLng(latitude, longitude);
		map?.setCenter(center);

		store.update((prev) => ({ ...prev, markers }));
	},
	clearMarkers() {
		const { markers } = get(store);

		markers.forEach((marker) => {
			marker.setMap(null);
		});

		store.update((prev) => ({ ...prev, markers: [] }));
	}
};

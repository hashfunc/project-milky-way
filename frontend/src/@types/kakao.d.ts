declare namespace kakao.maps {
	export class LatLng {
		constructor(latitude: string | number, longitude: string | number);
	}

	export class Map {
		constructor(container: HTMLElement, option: object);

		setCenter(center: LatLng): void;
	}

	export class Marker {
		constructor(option: object);

		setMap(map: Map | null): void;

		setPosition(position: LatLng): void;
	}

	export class MarkerImage {
		constructor(...option: object);
	}

	export class Size {
		constructor(width: number, height: number);
	}
}

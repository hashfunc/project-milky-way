declare namespace kakao.maps {
	export class LatLng {
		private readonly __nominal: void;

		constructor(latitude: string | number, longitude: string | number);
	}

	export class Map {
		private readonly __nominal: void;

		constructor(container: HTMLElement, option: object);

		setCenter(center: LatLng): void;
	}

	export class Marker {
		private readonly __nominal: void;

		constructor(option: object);

		setMap(map: Map | null): void;

		setPosition(position: LatLng): void;
	}

	export class MarkerImage {
		private readonly __nominal: void;

		constructor(...option: object);
	}

	export class Size {
		private readonly __nominal: void;

		constructor(width: number, height: number);
	}
}

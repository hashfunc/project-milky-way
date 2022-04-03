<script lang="ts">
	import { get } from 'svelte/store';
	import { store as mapStore } from '$lib/store/map';
	import IconStar from '$lib/assets/icon-star.png';
	import type { StarResponse } from './types';

	const DISTANCE_OPTIONS = [
		{ name: '500m', value: 500 },
		{ name: '1km', value: 1000 },
		{ name: '2km', value: 2000 },
		{ name: '3km', value: 3000 }
	];

	const URL = `${import.meta.env.VITE_API_BASE_URL}/api/v1/stars`;

	let distance = 1000;

	async function onSearch() {
		const { latitude, longitude } = get(mapStore);

		const query = `latitude=${latitude}&longitude=${longitude}&distance=${distance}`;

		try {
			const response = await fetch(`${URL}?${query}`);

			const data: StarResponse = await response.json();

			const { map, markers: prevMarkers } = get(mapStore);

			prevMarkers.forEach((marker) => {
				marker.setMap(null);
			});

			const markers = data.data.map((el) => {
				const image = new kakao.maps.MarkerImage(IconStar, new kakao.maps.Size(32, 32));
				const position = new kakao.maps.LatLng(el.latitude, el.longitude);
				return new kakao.maps.Marker({ map, position, image });
			});

			const center = new kakao.maps.LatLng(latitude, longitude);
			map?.setCenter(center);

			mapStore.update((prev) => ({ ...prev, markers }));
		} catch (error) {
			console.log(error);
		}
	}
</script>

<div class="search-control">
	<div class="distance-control">
		{#each DISTANCE_OPTIONS as option (option.value)}
			<div>
				<input id={`DO-${option.value}`} type="radio" value={option.value} bind:group={distance} />
				<label for={`DO-${option.value}`}>{option.name}</label>
			</div>
		{/each}
	</div>
	<button on:click={onSearch}>은하수 찾기</button>
</div>

<style lang="scss">
	.search-control {
		z-index: 10;
		position: fixed;
		bottom: 40px;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		box-sizing: border-box;
		padding: 12px;
		background-color: #ffffff;
		border: 1px solid #4f4f4f;
		border-radius: 4px;
		button {
			box-sizing: border-box;
			width: 180px;
			height: 40px;
			background-color: #efefef;
			border: 1px solid #4f4f4f;
			border-radius: 6px;
		}
	}
	.distance-control {
		display: flex;
		flex-direction: row;
		gap: 4px;
		box-sizing: border-box;
		margin-bottom: 12px;
		label {
			display: inline-block;
			box-sizing: border-box;
			padding: 4px;
			width: 65px;
			background-color: #efefef;
			border: 1px solid #4f4f4f;
			border-radius: 6px;
			text-align: center;
		}
		input[type='radio'] {
			display: none;
			&:checked + label {
				background-color: #3b8b5b;
				color: #efefef;
			}
		}
	}
</style>

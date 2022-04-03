<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import type { SearchData } from './types';

	export let result: SearchData[];

	const dispatch = createEventDispatcher();

	function select(latitude: string, longitude: string) {
		dispatch('select', { latitude, longitude });
	}
</script>

<div class="search-result">
	{#each result as item (item.id)}
		<div class="item" on:click={() => select(item.y, item.x)}>
			<div>{item.place_name}</div>
			<div>{item.road_address_name || item.address_name}</div>
		</div>
	{/each}
</div>

<style lang="scss">
	.search-result {
		overflow: auto;
		flex: 1;
		box-sizing: border-box;
		margin-top: 15px;
		width: 100%;
		border: 1px solid #efefef;
		.item {
			display: flex;
			justify-content: space-between;
			align-items: center;
			box-sizing: border-box;
			padding: 10px 15px;
			width: 100%;
			height: 45px;
			font-size: 12px;
			&:hover {
				background-color: #fafafa;
			}
		}
		.item + .item {
			border-top: 1px solid #efefef;
		}
	}
</style>

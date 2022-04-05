<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import NoResult from './NoResult.svelte';
	import SearchError from './SearchError.svelte';
	import SearchForm from './SearchForm.svelte';
	import SearchResult from './SearchResult.svelte';
	import { actions as mapActions } from '$lib/store/map';
	import type { SearchData, SearchResponse } from './types';

	const dispatch = createEventDispatcher();

	let contentDOM: HTMLElement;
	let result: SearchData[] | Error | null;

	function onClose(event: MouseEvent) {
		if (!contentDOM.contains(event.target as HTMLElement)) {
			dispatch('close');
		}
	}

	function onSearch(event: CustomEvent<SearchResponse | Error>) {
		if (event.detail instanceof Error) {
			result = event.detail;
		} else {
			result = event.detail?.data;
		}
	}

	function onSelect(event: CustomEvent) {
		mapActions.clearMarkers();
		mapActions.setCenter(event.detail);
		dispatch('close');
	}
</script>

<div class="search-modal" on:click={onClose}>
	<div class="content" bind:this={contentDOM}>
		<SearchForm on:search={onSearch} />
		{#if result instanceof Error}
			<SearchError />
		{:else if result && result.length}
			<SearchResult {result} on:select={onSelect} />
		{:else if result}
			<NoResult />
		{/if}
	</div>
</div>

<style lang="scss">
	.search-modal {
		z-index: 30;
		position: fixed;
		top: 0;
		left: 0;
		display: flex;
		justify-content: center;
		align-items: center;
		box-sizing: border-box;
		width: 100vw;
		height: 100vh;
		background-color: rgba(0, 0, 0, 0.2);
		.content {
			display: flex;
			flex-direction: column;
			box-sizing: border-box;
			padding: 30px 20px 20px 20px;
			width: 600px;
			height: 400px;
			background-color: #fff;
		}
	}
</style>

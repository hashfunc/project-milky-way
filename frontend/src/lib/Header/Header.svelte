<script lang="ts">
	import ProjectTitle from './ProjectTitle.svelte';
	import SearchButton from './SearchButton.svelte';
	import SearchModal from '$lib/SearchModal/SearchModal.svelte';

	let isSearchModalOpen: boolean = false;

	function openSearchModal() {
		isSearchModalOpen = true;
	}

	function closeSearchModal() {
		isSearchModalOpen = false;
	}
</script>

<header>
	<ProjectTitle />
	<SearchButton on:click={openSearchModal} />
	{#if isSearchModalOpen}
		<SearchModal on:close={closeSearchModal} />
	{/if}
</header>

<style>
	header {
		z-index: 10;
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		box-sizing: border-box;
		padding: 0 25px;
		height: 65px;
		background-color: #fff;
	}
</style>

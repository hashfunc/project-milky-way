<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	const URL = `${import.meta.env.VITE_API_BASE_URL}/api/v1/search/keyword`;

	let keyword: string = '';

	async function onSubmit() {
		const query = `keyword=${encodeURIComponent(keyword)}`;

		try {
			const response = await fetch(`${URL}?${query}`);

			if (response.status != 200) {
				throw new Error('Failed to search');
			}

			dispatch('search', await response.json());
		} catch (error) {
			dispatch('search', error);
		}
	}
</script>

<form class="search-form" on:submit|preventDefault={onSubmit}>
	<input type="text" placeholder="주소검색" bind:value={keyword} />
	<button>검색</button>
</form>

<style lang="scss">
	.search-form {
		display: flex;
		box-sizing: border-box;
		width: 100%;
		height: 35px;
		border: 1px solid #efefef;
		input {
			flex: 1;
			box-sizing: border-box;
			padding: 0 12px;
			height: 100%;
			background-color: #ffffff;
			border: none;
		}
		button {
			box-sizing: border-box;
			width: 50px;
			height: 100%;
			background-color: #ffffff;
			border: none;
			border-left: 1px solid #efefef;
		}
	}
</style>

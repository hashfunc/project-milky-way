<script lang="ts">
	import ProjectTitle from './ProjectTitle.svelte';
</script>

<header>
	<ProjectTitle />
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
		align-items: center;
		box-sizing: border-box;
		padding: 0 25px;
		height: 65px;
		background-color: #fff;
	}
</style>

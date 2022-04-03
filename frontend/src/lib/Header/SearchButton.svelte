<script>
	import IconSearch from '$lib/assets/icon-search.png';
</script>

<button class="search-button" type="button" on:click>
	<img src={IconSearch} alt="search icon" />
	주소검색
</button>

<style lang="scss">
	.search-button {
		box-sizing: border-box;
		width: 86px;
		height: 32px;
		border: 1px solid #efefef;
		background-color: #fefefe;
		img {
			box-sizing: border-box;
			margin-right: 2px;
			padding-bottom: 1px;
			width: 13px;
			vertical-align: middle;
		}
	}
</style>

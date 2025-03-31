<script lang="ts">
	import '../app.css';

	let { children } = $props();
	export const ssr = false;
</script>

{@render children()}

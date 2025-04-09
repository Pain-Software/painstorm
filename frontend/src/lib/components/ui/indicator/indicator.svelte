
<script lang="ts">
	import { type Props, indicatorVariants } from "./index.js";
	import { cn } from "$lib/utils.js";

	type $$Props = Props;

	let className: $$Props["class"] = undefined;
	export let variant: $$Props["variant"] = "default";
	export { className as class };
</script>

<div class="border px-4 py-2 text-card-foreground tracking-tight text-sm font-medium rounded-sm flex gap-4 items-center bg-card hover:bg-muted/50">
    <span class={cn(indicatorVariants({ variant, className }))}></span>
    <slot />
</div>

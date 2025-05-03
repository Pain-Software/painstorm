<script lang="ts">
	import { onMount, tick } from 'svelte';
	import { cn } from '$lib/utils.js';
	import { Button } from '$lib/components/ui/button';

	import * as Command from '$lib/components/ui/command';
	import * as Popover from '$lib/components/ui/popover';

	import {
		CheckIcon,
		ChevronsUpDown
	} from 'lucide-svelte';

    let { selected=$bindable(), modes } = $props();
	let open = $state(false);
	let className: string | undefined | null = undefined;

	function closeAndRefocusTrigger(triggerId: string) {
		open = false;

		tick().then(() => document.getElementById(triggerId)?.focus());
	}
</script>

<Popover.Root bind:open let:ids>
	<Popover.Trigger asChild let:builder>
		<Button
			builders={[builder]}
			variant="outline"
			role="combobox"
			aria-expanded={open}
			aria-label="Select a team"
			class={cn('w-[300px] justify-between', className)}
		>
			{@const CurrentIcon = selected.icon}
			<CurrentIcon class="mr-4 h-4 w-4" />
			{selected.label}
			<ChevronsUpDown class="ml-auto h-4 w-4 shrink-0 opacity-50" />
		</Button>
	</Popover.Trigger>
	<Popover.Content class="w-[300px] p-0">
		<Command.Root>
			<Command.List>
				{#each modes as mode}
					{@const Icon = mode.icon}
					<Command.Item
						onSelect={() => {
							selected = mode;
							closeAndRefocusTrigger(ids.trigger);
						}}
						value={mode.label}
						class="text-sm"
					>
						<Icon class="mr-4 h-4 w-4" />
						{mode.label}
						<CheckIcon
							class={cn('ml-auto h-4 w-4', selected.label !== mode.label && 'text-transparent')}
						/>
					</Command.Item>
				{/each}
			</Command.List>
		</Command.Root>
	</Popover.Content>
</Popover.Root>

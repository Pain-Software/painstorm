<script lang="ts">
	import { onMount } from 'svelte';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import { Button } from '$lib/components/ui/button';
	import { DatePickerWithRange } from '$lib/components/ui/datepicker';
	import { Input } from '$lib/components/ui/input';
	import { Indicator } from '$lib/components/ui/indicator';
	import * as Popover from '$lib/components/ui/popover';
	import ModeSwitcher from '$lib/fragments/mode-switcher.svelte';

	import PagePlaceCurrentStats from '$lib/fragments/page-place-current-stats.svelte';
	import PagePlaceHistoricStats from '$lib/fragments/page-place-historic-stats.svelte';
	import PageRainIntensity from '$lib/fragments/page-rain-intensity.svelte';
	import PageStableWeather from '$lib/fragments/page-stable-weather.svelte';
	import PageTempDiffs from "$lib/fragments/page-temp-diffs.svelte";

	import {
		CalendarIcon,
		CloudRainIcon,
		MapPinHouseIcon,
		ScaleIcon,
		ThermometerIcon
	} from 'lucide-svelte';
	import type { DateRange } from 'bits-ui';
	import { CalendarDate } from '@internationalized/date';
	import { Calendar } from '$lib/components/ui/calendar';
	import { cn } from 'tailwind-variants';

	const modes = [
		{
			label: 'Hledat aktuální data podle místa',
			value: 'place-current',
			icon: MapPinHouseIcon,
			page: PagePlaceCurrentStats
		},
		{
			label: 'Hledat historická data podle místa',
			value: 'place-historic',
			icon: MapPinHouseIcon,
			page: PagePlaceHistoricStats
		},
		{
			label: 'Hledat stabilní počasí',
			value: 'stable_weather',
			icon: ScaleIcon,
			page: PageStableWeather
		},
		{
			label: 'Hledat podle intenzity deště',
			value: 'intensity',
			icon: CloudRainIcon,
			page: PageRainIntensity
		},
		{
			label: 'Hledat rozdíly teplot',
			value: 'temp_diff',
			icon: ThermometerIcon,
			page: PageTempDiffs
		}
	];

	const now = new Date();

	let api: null | boolean = $state(null);
	let place = $state('');
	let intensity = $state(undefined);
	let mode = $state(modes[0]);
	let longtitude = $state();
	let latitude = $state();
	let count = $state();
	let weatherType = $state();

	let PageComponent = $derived(mode.page);

	let date: DateRange = $state({
		start: new CalendarDate(now.getFullYear(), now.getMonth() + 1, now.getDate()).subtract({
			days: 30
		}),
		end: new CalendarDate(now.getFullYear(), now.getMonth() + 1, now.getDate())
	});

	onMount(async () => {
		const route = new URL('/api/version', PUBLIC_BACKEND_URL);
		const response = await fetch(route);

		api = response.ok;
	});

	$inspect(date);
</script>

<main class="flex h-screen flex-col">
	<div class="flex h-[12vh] flex-col justify-center gap-2 border-b px-8">
		<div class="grid w-full grid-cols-3 items-center">
			<div class="justify-self-start">
				<ModeSwitcher bind:selected={mode} {modes} />
			</div>

			<div class="flex w-full justify-center">
				<h1 class="text-xl font-bold">Painstorm</h1>
			</div>

			<div class="flex w-full justify-end space-x-2">
				<Indicator variant={api == null ? 'default' : api ? 'online' : 'offline'}>API</Indicator>
			</div>
		</div>
		<div class="flex gap-2">
			{#if mode.value == 'place-current'}
				<Input
					type="text"
					placeholder="Název místa"
					class="h-9 md:w-[150px] lg:w-[200px]"
					bind:value={place}
				/>
				<Input
					type="number"
					placeholder="Zeměpisná šířka"
					class="h-9 md:w-[150px] lg:w-[200px]"
					bind:value={latitude}
				/>
				<Input
					type="number"
					placeholder="Zeměpisná délka"
					class="h-9 md:w-[150px] lg:w-[200px]"
					bind:value={longtitude}
				/>
				<Input
					type="number"
					placeholder="Počet dní"
					min={1}
					step={1}
					max={7}
					class="h-9 md:w-[150px] lg:w-[200px]"
					bind:value={count}
				/>
			{:else if mode.value == 'stable_weather'}
				<Input
					type="text"
					placeholder="Název místa"
					class="h-9 md:w-[150px] lg:w-[200px]"
					bind:value={place}
				/>
				<Input
					type="number"
					placeholder="Zeměpisná šířka"
					class="h-9 md:w-[150px] lg:w-[200px]"
					bind:value={latitude}
				/>
				<Input
					type="number"
					placeholder="Zeměpisná délka"
					class="h-9 md:w-[150px] lg:w-[200px]"
					bind:value={longtitude}
				/>
				<Input
					type="text"
					placeholder="Typ počasí"
					class="h-9 md:w-[150px] lg:w-[200px]"
					bind:value={weatherType}
				/>
				<DatePickerWithRange bind:value={date} />
			{:else if mode.value == 'temp_diff'}
				<Popover.Root>
					<Popover.Trigger>
						<Button
							variant="outline"
							class={cn(
								'w-[280px] justify-start text-left font-normal',
								!date.end && 'text-muted-foreground'
							)}
						>
							<CalendarIcon class="mr-2 h-4 w-4" />
							{date.end ? new Date(date.end.toString()).toLocaleDateString() : 'Pick a date'}
						</Button>
					</Popover.Trigger>
					<Popover.Content class="w-auto p-0">
						<Calendar bind:value={date.end} initialFocus />
					</Popover.Content>
				</Popover.Root>
			{:else if mode.value == 'place-historic'}
				<Input
					type="text"
					placeholder="Název místa"
					class="h-9 md:w-[150px] lg:w-[200px]"
					bind:value={place}
				/>
				<Input
					type="number"
					placeholder="Zeměpisná šířka"
					class="h-9 md:w-[150px] lg:w-[200px]"
					bind:value={latitude}
				/>
				<Input
					type="number"
					placeholder="Zeměpisná délka"
					class="h-9 md:w-[150px] lg:w-[200px]"
					bind:value={longtitude}
				/>
				<DatePickerWithRange bind:value={date} />
			{:else if mode.value == 'intensity'}
				<Input
					type="number"
					placeholder="Intenzita deště"
					class="h-9 md:w-[150px] lg:w-[200px]"
					bind:value={intensity}
				/>
				<DatePickerWithRange bind:value={date} />
			{/if}
		</div>
	</div>

	<div class="h-[88vh] flex-1">
		<PageComponent {weatherType} {count} {place} {intensity} from={date.start} to={date.end} />
	</div>
</main>

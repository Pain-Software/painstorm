<script lang="ts">
	import _ from 'lodash';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import * as Table from '$lib/components/ui/table';
	import { ScrollArea } from '$lib/components/ui/scroll-area';
	import type { ComponentType } from 'svelte';
	import * as Card from '$lib/components/ui/card';
	import { Skeleton } from '$lib/components/ui/skeleton';

	const { place, from, to, intensity, count } = $props();

	interface Measurement {
		id: number;
		timestamp: string;
		city_id: number;
		min_temperature: number;
		max_temperature: number;
		temperature: number;
		humidity: number;
		pressure: number;
        sea_level?: number;
        ground_level?: number;
		wind_speed: number;
		wind_degrees: number;
		weather: {
			id: number;
			main: string;
			description: string;
			icon: string;
		}[];
        clouds?: number;
        cnt?: number;
	}

	interface Statistic {
		name: string;
		value: string;
		description: string;
		Icon: ComponentType;
	}

	let history = $state<Measurement[]>([]);
	let stats = $state<Statistic[]>([]);
	let isLoading = $state(true);

	$effect(() => {
		void place;
		void from;
		void to;

		if (!place || !from || !to) return;

		async function getHistoryData() {
			const url = new URL('/api/weather/search/retrieve', PUBLIC_BACKEND_URL);
			url.searchParams.set('place', String(place));
			url.searchParams.set('from', String(new Date(from).valueOf() / 1000));
			url.searchParams.set('to', String(new Date(to).valueOf() / 1000));

			const response = await fetch(url);
			const data = await response.json();

			if (!Array.isArray(data)) {
                history = [];
                return;
            }
				
            history = data.toSorted((a: Measurement, b: Measurement) =>
                Number(new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime())
            );

            isLoading = false;
		}

		const getHistoryDataDebounced = _.debounce(getHistoryData, 250);

		getHistoryDataDebounced();

		// Optional: cancel on teardown (e.g., if component is destroyed)
		return () => {
			getHistoryDataDebounced.cancel();
		};
	});
</script>

{#if !place}
	<div class="grid h-full w-full place-items-center">
		<h1 class="text-muted-foreground text-xl">Pro zobrazení statistik zadejte název místa</h1>
	</div>
{:else}
	<div class="hidden flex-col md:flex">
		<div class="flex-1 p-8">
			<div class="flex flex-col">
				<h2 class="text-3xl font-bold tracking-tight">
					Měření v {place}
				</h2>
                <h3 class="text-xl text-muted-foreground">
                    Od {new Date(from.toString()).toLocaleDateString()} do {new Date(to.toString()).toLocaleDateString()}
                </h3>
			</div>
		</div>

		<div class="flex flex-col gap-4 p-8 pt-0">
			<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
				{#each stats as { name, value, description, Icon }}
					{#if isLoading}
						<Card.Root>
							<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
								<Card.Title><Skeleton class="h-3.5 w-32" /></Card.Title>
								<Skeleton class="h-4 w-4" />
							</Card.Header>
							<Card.Content class="flex flex-col gap-1">
								<Skeleton class="h-8 w-64" />
								<Skeleton class="h-4 w-72" />
							</Card.Content>
						</Card.Root>
					{:else}
						<Card.Root>
							<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
								<Card.Title class="text-sm font-medium">{name}</Card.Title>
								<Icon class="text-muted-foreground h-4 w-4" />
							</Card.Header>
							<Card.Content>
								<div class="text-2xl font-bold">{value}</div>
								<p class="text-muted-foreground text-xs">{description}</p>
							</Card.Content>
						</Card.Root>
					{/if}
				{/each}
			</div>

			<ScrollArea class="h-[60vh] rounded-md border p-4">
				<Table.Root>
					<Table.Header>
						<Table.Row>
							<Table.Head>Počasí</Table.Head>
							<Table.Head>Časové razítko</Table.Head>
							<Table.Head>Průměrná teplota</Table.Head>
							<Table.Head>Minimální teplota</Table.Head>
							<Table.Head>Maximální teplota</Table.Head>
							<Table.Head>Vlhkost</Table.Head>
							<Table.Head>Atmosférický tlak</Table.Head>
							<Table.Head>Rychlost větru</Table.Head>
							<Table.Head>Směr větru</Table.Head>
						</Table.Row>
					</Table.Header>
					<Table.Body>
						{#each history as row}
							<Table.Row>
								<Table.Cell class="flex items-center gap-2">
									<img
										src={`https://openweathermap.org/img/wn/${row.weather[0].icon}@2x.png`}
										alt=""
										class="h-12 w-12"
									/>
									<span>{row.weather[0].main}</span>
								</Table.Cell>
								<Table.Cell class="font-medium"
									>{new Date(row.timestamp).toLocaleString()}</Table.Cell
								>
								<Table.Cell>{row.temperature} °C</Table.Cell>
								<Table.Cell>{row.min_temperature} °C</Table.Cell>
								<Table.Cell>{row.max_temperature} °C</Table.Cell>
								<Table.Cell>{row.humidity} %</Table.Cell>
								<Table.Cell>{row.pressure} hPA</Table.Cell>
								<Table.Cell>{row.wind_speed} m/s</Table.Cell>
								<Table.Cell>{row.wind_degrees ? `${row.wind_degrees} °` : ""}</Table.Cell>
							</Table.Row>
						{/each}
					</Table.Body>
				</Table.Root>
			</ScrollArea>
		</div>
	</div>
{/if}

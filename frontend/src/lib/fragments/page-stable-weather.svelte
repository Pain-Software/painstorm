<script lang="ts">
    import _ from 'lodash';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import * as Table from '$lib/components/ui/table';
	import { ScrollArea } from '$lib/components/ui/scroll-area';

	const { place, longtitude, latitude, from, to, weatherType } = $props();

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

	let history = $state<Measurement[][]>([]);
	let isLoading = $state(true);

    let noPlaceDefined = $derived(!place && (!longtitude || !latitude));

	$effect(() => {
		void place;
        void longtitude;
		void latitude;
		void from;
		void to;

		if (noPlaceDefined || !from || !to || !weatherType) return;

		async function getStableWeather() {
			const url = new URL('/api/weather/search/stable-weather', PUBLIC_BACKEND_URL);
			url.searchParams.set('place', String(place));
            if (longtitude) url.searchParams.set('lon', String(longtitude))
            if (latitude) url.searchParams.set('lat', String(latitude))
			url.searchParams.set('from', String(from.toString()));
			url.searchParams.set('to', String(to.toString()));
            url.searchParams.set('weatherType', String(weatherType));

			const response = await fetch(url);
			const data = await response.json();

			if (!Array.isArray(data)) {
                history = [];
                return;
            }
				
            history = _.compact(data);
            isLoading = false;
		}

		const getStableWeatherDebounced = _.debounce(getStableWeather, 250);

		getStableWeatherDebounced();

		// Optional: cancel on teardown (e.g., if component is destroyed)
		return () => {
			getStableWeatherDebounced.cancel();
		};
	});
</script>

{#if noPlaceDefined}
	<div class="grid h-full w-full place-items-center">
		<h1 class="text-muted-foreground text-xl">Pro zobrazení statistik zadejte název místa.</h1>
	</div>
{:else if history.length == 0}
    <div class="grid h-full w-full place-items-center">
        <h1 class="text-muted-foreground text-xl">Pro dané místo nebyly nalezeny žádné výsledky.</h1>
    </div>
{:else}
	<div class="hidden flex-col md:flex">
		<div class="flex-1 p-8">
			<div class="flex flex-col">
				<h2 class="text-3xl font-bold tracking-tight">
					Stabilní počasí v {place}
				</h2>
                <h3 class="text-xl text-muted-foreground">
                    Od {new Date(from.toString()).toLocaleDateString()} do {new Date(to.toString()).toLocaleDateString()}
                </h3>
			</div>
		</div>

		<div class="flex flex-col gap-4 p-8 pt-0">
			<ScrollArea class="h-[60vh] rounded-md border p-4">
				<Table.Root>
					<Table.Header>
						<Table.Row>
							<Table.Head>Počasí</Table.Head>
                            <Table.Head>Oblačnost</Table.Head>
							<Table.Head>Začátek</Table.Head>
							<Table.Head>Konec</Table.Head>
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
										src={`https://openweathermap.org/img/wn/${row[0].weather[0].icon}@2x.png`}
										alt=""
										class="h-12 w-12"
									/>
									<span>{row[0].weather[0].main}</span>
								</Table.Cell>
                                <Table.Cell>{row[0].clouds} %</Table.Cell>
								<Table.Cell>{new Date(row[0].timestamp).toLocaleDateString()}</Table.Cell>
								<Table.Cell>{new Date(row[row.length - 1].timestamp).toLocaleDateString()}</Table.Cell>
                                <Table.Cell>{row[0].temperature} °C</Table.Cell>
								<Table.Cell>{row[0].min_temperature} °C</Table.Cell>
								<Table.Cell>{row[0].max_temperature} °C</Table.Cell>
								<Table.Cell>{row[0].humidity} %</Table.Cell>
								<Table.Cell>{row[0].pressure} hPA</Table.Cell>
								<Table.Cell>{row[0].wind_speed} m/s</Table.Cell>
								<Table.Cell>{row[0].wind_degrees ? `${row[0].wind_degrees} °` : ""}</Table.Cell>
							</Table.Row>
						{/each}
					</Table.Body>
				</Table.Root>
			</ScrollArea>
		</div>
	</div>
{/if}

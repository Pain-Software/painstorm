<script lang="ts">
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import * as Table from '$lib/components/ui/table';

	const { from, to, intensity } = $props();

	interface ResponseBody {
		id: number;
		name: string;
		find_name: string;
		country: string;
		longitude: number;
		latitude: number;
	}

	let rows = $state<ResponseBody[]>([]);
	const formattedFrom = $derived(new Date(from).toLocaleDateString());
	const formattedTo = $derived(new Date(to).toLocaleDateString());

    
	$effect(() => {
		if (!from || !to || intensity == null) return;
        
        const url = new URL('/api/weather/search/intensity', PUBLIC_BACKEND_URL);
        url.searchParams.set('intensity', String(intensity));
        url.searchParams.set('from', String(new Date(from).valueOf() / 1000));
        url.searchParams.set('to', String(new Date(to).valueOf() / 1000));
    
		(async () => {
			const response = await fetch(url);
			const data = await response.json();
			rows = data;
		})();
	});
</script>

{#if intensity == undefined}
	<div class="grid h-full w-full place-items-center">
		<h1 class="text-muted-foreground text-xl">Pro zobrazení statistik zadejte intenzitu deště</h1>
	</div>
{:else}
	<div class="hidden flex-col md:flex">
		<div class="flex-1 space-y-4 p-8 pt-6">
			<div class="flex flex-col">
				<h2 class="text-3xl font-bold tracking-tight">
					Intenzita deště 
				</h2>
                <h3 class="text-xl text-muted-foreground">
                    od {formattedFrom} do {formattedTo}
                </h3>
			</div>
		</div>

		<div class="p-8">
			<Table.Root>
				<Table.Header>
					<Table.Row>
						<Table.Head class="w-[100px]">Město</Table.Head>
						<Table.Head>Find Name</Table.Head>
						<Table.Head>Country</Table.Head>
                        <Table.Head>Zeměpisná délka</Table.Head>
						<Table.Head>Zeměpisná šířka</Table.Head>
					</Table.Row>
				</Table.Header>
				<Table.Body>
					{#each rows as row}
						<Table.Row>
							<Table.Cell class="font-medium">{row.name}</Table.Cell>
							<Table.Cell>{row.find_name}</Table.Cell>
							<Table.Cell>{row.country}</Table.Cell>
							<Table.Cell>{row.longitude}</Table.Cell>
                            <Table.Cell>{row.latitude}</Table.Cell>
						</Table.Row>
					{/each}
				</Table.Body>
			</Table.Root>
		</div>
	</div>
{/if}

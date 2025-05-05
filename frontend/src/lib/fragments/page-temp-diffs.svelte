<script lang="ts">
	import { PUBLIC_BACKEND_URL } from "$env/static/public";
    import * as Card from "$lib/components/ui/card"
    import * as Table from "$lib/components/ui/table"
	import _ from "lodash";

    let { to } = $props();

    interface TempDiffResponse {
        id: number;
        name: string; 
        find_name: string;
        country: string;
        longitude: number
        latitude: number
        tempDiff: number 
    }

    let temps = $state<TempDiffResponse[]>([]);
    
    $effect(() => {
		void to;

        console.warn(to);
		if (!to) return;

		async function getTempDiffs() {
			const url = new URL('/api/weather/search/temp-diff', PUBLIC_BACKEND_URL);
			url.searchParams.set('date', String(to.toString()));

			const response = await fetch(url);
			const data = await response.json();

			if (!Array.isArray(data)) {
                temps = [];
                return;
            }
            
            temps = data;
		}

		const getTempDiffsDebounced = _.debounce(getTempDiffs, 250);

		getTempDiffsDebounced();

		// Optional: cancel on teardown (e.g., if component is destroyed)
		return () => {
			getTempDiffsDebounced.cancel();
		};
	});

</script>

<div class="flex flex-col gap-4 p-8">
    <Card.Root class="col-span-3">
        <Card.Header>
            <Card.Title>Místa s největším rozdílem teplot</Card.Title>
            <Card.Description>You made 265 sales this month.</Card.Description>
        </Card.Header>
        <Card.Content>
            <Table.Root>
                <Table.Header>
                    <Table.Row>
                        <Table.Head>Místo</Table.Head>
                        <Table.Head>Stát</Table.Head>
                        <Table.Head>Teplotní rozdíl</Table.Head>
                    </Table.Row>
                </Table.Header>
                <Table.Body>
                    {#each temps as row}
                        <Table.Row>
                            <Table.Cell class="font-medium">{row.name}</Table.Cell>
                            <Table.Cell>{row.country}</Table.Cell>
                            <Table.Cell>{Number(row.tempDiff).toLocaleString("cs-CZ", { type: "decimal", maximumFractionDigits: 2, minimumFractionDigits: 2 })} °C</Table.Cell>
                        </Table.Row>
                    {/each}
                </Table.Body>
            </Table.Root>
        </Card.Content>
    </Card.Root>
</div>

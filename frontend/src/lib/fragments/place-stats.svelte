<script lang="ts">
	import * as Card from '$lib/components/ui/card';
    import { onMount, type ComponentType } from 'svelte';
    import { Skeleton } from '$lib/components/ui/skeleton';
	import {
		Globe,
		Waves,
		Mountain,
		Droplet,
		Gauge,
		Milestone,
		ThermometerSnowflake,
		ThermometerSun
	} from 'lucide-svelte';

    interface Statistic {
        name: string;
        value: string;
        description: string
        Icon: ComponentType
    }

    let isLoading = $state<boolean>(true);
    let stats = $state<Statistic[]>(new Array(8).fill({name: "", value: "", description: "", Icon: Globe}));

    async function getData() {
		// We bangin'
		await new Promise((resolve) => setTimeout(resolve, 5000))
        const numberFormatter = new Intl.NumberFormat('cs-CZ', { style: 'decimal' });
        
        stats = [
            {
                name: "Atmosférický tlak",
                value: `${numberFormatter.format(1021)} hPa`,
                description: new Date().toLocaleString(),
                Icon: Globe
            },
            {
                name: "Atmosférický tlak u hladiny moře",
                value: `${numberFormatter.format(1021)} hPa`,
                description: new Date().toLocaleString(),
                Icon: Waves
            },
            {
                name: "Atmosférický tlak u země",
                value: `${numberFormatter.format(991)} hPa`,
                description: new Date().toLocaleString(),
                Icon: Mountain
            },
            {
                name: "Vlhkost",
                value: `${Number(60 / 100).toLocaleString("cs-CZ", { style: "percent" })}`,
                description: new Date().toLocaleString(),
                Icon: Droplet
            },
            {
                name: "Rychlost větru",
                value: `${numberFormatter.format(69)} m/s`,
                description: new Date().toLocaleString(),
                Icon: Gauge
            },
            {
                name: "Směr větru",
                value: `${numberFormatter.format(121)} °`,
                description: new Date().toLocaleString(),
                Icon: Milestone
            },
            {
                name: "Minimální teplota",
                value: `${numberFormatter.format(31)} °C`,
                description: new Date().toLocaleString(),
                Icon: ThermometerSnowflake
            },
            {
                name: "Maximální teplota",
                value: `${numberFormatter.format(69)} °C`,
                description: new Date().toLocaleString(),
                Icon: ThermometerSun
            }
        ];

		isLoading = false;
	}

	onMount(() => {
		getData();
	});
</script>

<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
    {#each stats as { name, value, description, Icon }}
        {#if isLoading}
            <Card.Root>
                <Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
                    <Card.Title><Skeleton class="w-32 h-3.5"/></Card.Title>
                    <Skeleton class="w-4 h-4"/>
                </Card.Header>
                <Card.Content class="flex flex-col gap-1">
                    <Skeleton class="w-64 h-8"/>
                    <Skeleton class="w-72 h-4"/>
                </Card.Content>
            </Card.Root>
        {:else}
            <Card.Root>
                <Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
                    <Card.Title class="text-sm font-medium">{name}</Card.Title>
                    <Icon class="text-muted-foreground h-4 w-4"/>
                </Card.Header>
                <Card.Content>
                    <div class="text-2xl font-bold">{value}</div>
                    <p class="text-muted-foreground text-xs">{description}</p>
                </Card.Content>
            </Card.Root>
        {/if}
    {/each}
</div>

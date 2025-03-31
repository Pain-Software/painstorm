<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import { Skeleton } from '$lib/components/ui/skeleton';
    import { onMount, type ComponentType } from 'svelte';
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
    let stats = $state<Statistic[]>([]);

    async function getData() {
		// We bangin'
		setTimeout(() => {}, 1000);

        const numberFormatter = new Intl.NumberFormat('cs-CZ', { style: 'decimal' });

        stats.push(
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
        );

		/*pressure.atmosphere.value = `${numberFormatter.format(1021)} hPa`;
		pressure.atmosphere.date = new Date().toLocaleString();

		pressure.sea.value = `${numberFormatter.format(1021)} hPa`;
		pressure.sea.date = new Date().toLocaleString();

		pressure.ground.value = `${numberFormatter.format(991)} hPa`;
		pressure.ground.date = new Date().toLocaleString();

        humidity.value = `${Number(60 / 100).toLocaleString("cs-CZ", { style: "percent" })}`;
        humidity.date = new Date().toLocaleString();

		wind.speed.value = `${numberFormatter.format(69)} m/s`;
		wind.speed.date = new Date().toLocaleString();

        wind.degree.value = `${numberFormatter.format(121)} °`;
		wind.degree.date = new Date().toLocaleString();

		temperature.max.value = `${numberFormatter.format(69)} °C`;
		temperature.max.date = new Date().toLocaleString();

		temperature.min.value = `${numberFormatter.format(10)} °C`;
		temperature.min.date = new Date().toLocaleString();*/

		isLoading = false;
	}

	onMount(() => {
		getData();
	});
</script>

<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
    {#each stats as { name, value, description, Icon }}
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
    {/each}
</div>

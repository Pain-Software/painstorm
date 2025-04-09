<script lang="ts">
	import Download from 'lucide-svelte/icons/download';

	import { Button } from '$lib/components/ui/button';
	import { DatePickerWithRange } from '$lib/components/ui/datepicker';
    import PlaceStats from '$lib/fragments/place-stats.svelte';
    import PlacesTempDiffs from '$lib/fragments/places-temp-diffs.svelte';
    import PlacesWithRainIntensity from '$lib/fragments/places-with-rain-intensity.svelte';
	import { Input } from '$lib/components/ui/input';
	
    
    let place = 'Testovací název místa';


</script>

<div class="flex justify-between items-center px-8 border-b h-16">
    <div class="flex items-center space-x-2">
        <Input type="search" placeholder="Název místa..." class="h-9 md:w-[100px] lg:w-[300px]" />
        <Input type="search" placeholder="Zeměpisná šířka" class="h-9 md:w-[50px] lg:w-[100px]" />
        <Input type="search" placeholder="Zeměpisná délka" class="h-9 md:w-[50px] lg:w-[100px]" />        
    </div>

    <div class="flex items-center space-x-2">
        <DatePickerWithRange />
        <Button size="sm">
            <Download class="mr-2 h-4 w-4" />
            Hledat
        </Button>
    </div>
</div>

<div class="hidden flex-col md:flex">
	<div class="flex-1 space-y-4 p-8 pt-6">
		<div class="flex items-center justify-between space-y-2">
			<h2 class="text-3xl font-bold tracking-tight">{place}</h2>
			<div class="flex items-center space-x-2">
				<Button size="sm">
					<Download class="mr-2 h-4 w-4" />
					Download
				</Button>
			</div>
		</div>
		
        <PlaceStats/>

		<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-9">
            <div class="col-span-3">
                <PlacesTempDiffs/>                
            </div>
            <div class="col-span-3">
                <PlacesWithRainIntensity value={69}/>
            </div>
		</div>
	</div>
</div>

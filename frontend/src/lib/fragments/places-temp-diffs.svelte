<script lang="ts">
    import * as Card from "$lib/components/ui/card"
    import * as Table from "$lib/components/ui/table"

    const temps = $state([
        {
            place: "Město 1",
            min: 10,
            max: 60
        },
        {
            place: "Město 2",
            min: 10,
            max: 60
        },
        {
            place: "Město 3",
            min: 10,
            max: 60
        },
        {
            place: "Město 4",
            min: 10,
            max: 60
        },
        {
            place: "Město 5",
            min: 10,
            max: 60
        },
        {
            place: "Město 6",
            min: 10,
            max: 60
        },
        {
            place: "Město 7",
            min: 10,
            max: 60
        },
        {
            place: "Město 8",
            min: 10,
            max: 60
        },
        {
            place: "Město 9",
            min: 10,
            max: 60
        },
        {
            place: "Město 10",
            min: 10,
            max: 60
        }
    ]);
</script>

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
                    <Table.Head>Minimální teplota</Table.Head>
                    <Table.Head>Maximální teplota</Table.Head>
                    <Table.Head>Teplotní rozdíl</Table.Head>
                </Table.Row>
            </Table.Header>
            <Table.Body>
                {#each temps as measurement}
                    <Table.Row>
                        <Table.Cell class="font-medium">{measurement.place}</Table.Cell>
                        <Table.Cell>{measurement.min}</Table.Cell>
                        <Table.Cell>{measurement.max}</Table.Cell>
                        <Table.Cell>{measurement.max - measurement.min}</Table.Cell>
                    </Table.Row>
                {/each}
            </Table.Body>
        </Table.Root>
    </Card.Content>
</Card.Root>

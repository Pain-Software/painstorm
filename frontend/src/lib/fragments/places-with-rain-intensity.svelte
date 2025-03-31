<script lang="ts">
    import * as Card from "$lib/components/ui/card"
    import * as Table from "$lib/components/ui/table"

    const {
        value
    } = $props();

    let places = $state([
        {
            id: 1,
            name: "Město 1"
        },
        {
            id: 2,
            name: "Město 2"
        },
        {
            id: 3,
            name: "Město 3"
        },
        {
            id: 4,
            name: "Město 4"
        },
        {
            id: 5,
            name: "Město 5"
        }
    ]);
</script>

<Card.Root class="col-span-3">
    <Card.Header>
        <Card.Title>Vlhkost vzduchu {value} %</Card.Title>
        <Card.Description>Seznam míst se zadanou vhlkostí vzduchu</Card.Description>
    </Card.Header>
    <Card.Content>
        <Table.Root>
            <Table.Header>
                <Table.Row>
                    <Table.Head>Místo</Table.Head>
                </Table.Row>
            </Table.Header>
            <Table.Body>
                {#each places as place}
                    <Table.Row>
                        <Table.Cell class="font-medium">{place.name}</Table.Cell>
                    </Table.Row>
                {/each}
            </Table.Body>
        </Table.Root>
    </Card.Content>
</Card.Root>

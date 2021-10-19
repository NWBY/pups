<template>
    <div style="text-align: left;">
        <h2>Volumes</h2>
        <table class="table" v-if="volumes.length">
            <thead>
                <th>Image</th>
                <th>State</th>
                <th>Status</th>
            </thead>
            <tbody>
                <item v-for="(volume, index) in volumes" :key="index" :item="volume"></item>
            </tbody>
        </table>
        <p v-else>No volumes</p>
    </div>
</template>

<script>
import Item from './Item.vue'

export default {
    props: {
        volumes: {
            type: Array,
            required: true
        }
    },
    components: {
        Item
    }
}
</script>

<style scoped>
.table {
    width: 100%;
    table-layout: fixed;
    text-align: left;
}
</style>
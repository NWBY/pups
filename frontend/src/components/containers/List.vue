<template>
    <div>
        <table class="table">
            <thead>
                <th>Image</th>
                <th>State</th>
                <th>Status</th>
            </thead>
            <tbody>
                <item v-for="(container, index) in containers" :key="index" :item="container"></item>
            </tbody>
        </table>
        
    </div>
</template>

<script>
import Item from './Item.vue'

export default {
    props: {
        containers: {
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
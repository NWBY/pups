<template>
    <tr>
        <td>{{item.Image}}</td>
        <td>{{item.State}}</td>
        <td>{{item.Status}}</td>
    </tr>
</template>

<script>
export default {
    props: {
        item: {
            type: Object,
            required: true
        }
    }
}
</script>
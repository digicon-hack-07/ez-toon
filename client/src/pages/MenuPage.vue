<template>
    <div :class="$style.menu">
        <thumbnail/>
        <thumbnail/>
        <thumbnail/>
        <thumbnail/>
        <thumbnail/>
        <thumbnail/>
        <thumbnail/>
        <thumbnail/>
        <thumbnail/>
        <thumbnail/>
        <thumbnail/>
    </div>
</template>

<script setup lang="ts">
import Thumbnail from '../components/menu/Thumbnail.vue'
</script>

<style module>
    .menu{
        text-align: left;
    }
</style>


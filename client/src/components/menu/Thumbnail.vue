<!-- eslint-disable vue/multi-word-component-names -->
<template>
    <div :class="$style.thumbnail">
        <router-link :to=url>
            <div>
                <img src="/vite.svg" :class=$style.image />
                <p>{{name}}</p>
            </div>
            
        </router-link>
    </div>
</template>

<script lang="ts" setup>
withDefaults(
  defineProps<{
    name: string
    url: string
    image: string
  }>(),
  {
    name: "noname",
    url: "/",
    image: "/vite.svg",
  }
)

</script>

<style module>

    .image {
        width : 8rem;
        height : 6rem;
    }
    

    .thumbnail {
        display: inline-block;
        text-align: center;
        border-radius: 0.5rem;
        padding: 1rem;
        margin: 2rem;
    }

    .thumbnail:hover {
        background-color:darkgray;
    }
</style>
<script setup lang="ts">
    import { computed, nextTick, onMounted, ref, useTemplateRef, watch } from 'vue';

    const fileName = ref("rythmic/memorex-memories-in-motion.mp3")
    const fileURL = computed(() => "/stream/" + fileName.value)
</script>
<template>
    <div class="container">

        <input v-model="fileName" class="name" />
        <audio ref="player" controls :src="fileURL" />
    </div>
</template>

<style scoped>
.container {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
    display: flex;
    flex-direction: column;
    align-items: center;
    width: min-content;
}

.name {
    width: 100%;
}
</style>
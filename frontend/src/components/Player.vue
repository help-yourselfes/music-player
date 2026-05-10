<script setup lang="ts">
    import { useAppStore } from '@/stores/app';
import { computed } from 'vue';
    const store = useAppStore()
    import {EventsOn} from "@runtime"
    import {Track} from "@/types"

    // const src = computed(() => "/files/static/stream/?path="+btoa(store.trackPath))
    const src = computed(() => "/files/static/stream/?path="+(store.trackPath))

    EventsOn("play:track", (track: Track) => {
        store.playTrack(track)
    })

</script>
<template>
    <div class="container">
        <audio ref="player" controls :src />
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
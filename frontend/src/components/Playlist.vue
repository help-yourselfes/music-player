<script setup lang="ts">
    import { computed, ref, watch } from 'vue';
    import { main } from "@go/models"
    import { GetTracks } from "@go/main/App"

    const path = ref("D:/media/music")
    const tracklist = ref<main.Track[]>([])

    watch(path, async () => {
        try {
            const res = await GetTracks(path.value)
            if (!res.Tracks) throw new Error("no tracks")
            tracklist.value = res.Tracks
        } catch (e) {
            console.error(e)
        }
    })
</script>
<template>
    <div class="container">
        <button v-for="track in tracklist" class="track">
            {{ track.name }}
        </button>
    </div>
</template>
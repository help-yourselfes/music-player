<script setup lang="ts">
    import { onMounted, ref, watch } from 'vue';
    import { main } from "@go/models"
    import { GetTracks } from "@go/main/App"
    import { useAppStore } from '@/stores/app';

    const path = ref("D:/media/music")
    const tracklist = ref<main.Track[]>([])

    const reloadTracks = async () => {
        try {
            const res = await GetTracks(path.value)
            console.log(res)
            if (!res.Tracks) throw new Error("no tracks")
            tracklist.value = res.Tracks
        } catch (e) {
            console.error(e)
        }
    }

    watch(path, reloadTracks)

    onMounted(() => {
        store.playlistPath = path.value
        reloadTracks()
    })

    const store = useAppStore()
</script>
<template>
    <div class="container">
        <button v-for="track in tracklist" @click="store.playTrack(track)" class="track">
            {{ track.name }}
        </button>
    </div>
</template>
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useAppStore = defineStore('app', () => {
    const playlistPath = ref("")
    const track = ref("")

    const trackPath = computed(() => `${playlistPath.value}/${track.value}`)

    return { playlistPath, track, trackPath }
})

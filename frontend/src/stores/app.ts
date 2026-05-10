import { Track } from '@/types'
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useAppStore = defineStore('app', () => {
    const playlistPath = ref("")
    const track = ref("")

    const trackPath = ref("")

    const playTrack = (t: Track) => {
        track.value = t.name
        trackPath.value = t.path
    }

    return { playlistPath, track, trackPath, playTrack }
})

<script setup lang="ts">
import { onMounted, ref } from "vue";
import L from "leaflet";

// 🔥 Fix missing marker icons in Vite/Nuxt
import markerIcon2x from "leaflet/dist/images/marker-icon-2x.png";
import markerIcon from "leaflet/dist/images/marker-icon.png";
import markerShadow from "leaflet/dist/images/marker-shadow.png";

delete (L.Icon.Default.prototype as any)._getIconUrl;

L.Icon.Default.mergeOptions({
  iconRetinaUrl: markerIcon2x,
  iconUrl: markerIcon,
  shadowUrl: markerShadow,
});

const props = defineProps<{
  lat: number;
  lng: number;
  label: string;
}>();

const mapEl = ref<HTMLDivElement | null>(null);

onMounted(() => {
  if (!mapEl.value) return;

  const map = L.map(mapEl.value).setView([props.lat, props.lng], 15);

  L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
    attribution: "&copy; OpenStreetMap contributors",
  }).addTo(map);

  L.marker([props.lat, props.lng])
    .addTo(map)
    .bindPopup(props.label)
    .openPopup();
});
</script>

<template>
  <div ref="mapEl" class="w-full h-full" />
</template>

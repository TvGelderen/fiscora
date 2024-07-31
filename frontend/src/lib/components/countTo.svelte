<script lang="ts">
	import { getFormattedAmount } from "$lib";

	const { start, value }: { start: number; value: number } = $props();

	let count = $state(start);

	const duration = 2000;
	const frames = (duration / 1000) * 60;

	$effect(() => {
		const delta = value - start;
		const delay = frames / Math.abs(delta);
		let step = delta / frames;

		const interval = setInterval(() => {
			if (
				(delta > 0 && count >= value) ||
				(delta < 0 && count <= value)
			) {
				if (count !== value) {
					count = value;
				}
				clear();
				return;
			}

			count += step;
		}, delay);

		const clear = () => clearInterval(interval);
	});
</script>

<span>{getFormattedAmount(Math.round(count))}</span>

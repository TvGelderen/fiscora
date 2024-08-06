<script lang="ts">
	import { getFormattedAmount } from "$lib";

	const { start, value }: { start: number; value: number } = $props();

	let count = $state(start);

	const duration = 1000;
	const frames = (duration / 1000) * 60;

	$effect(() => {
		const delta = value - start;
		if (delta === 0) return;
		const delay = frames / Math.abs(delta);
		const step =
			delta > 0
				? Math.max(delta / frames, 0.5)
				: Math.min(delta / frames, -0.5);

		console.log(`From ${start} to ${value} with ${step}`);

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

<span>{getFormattedAmount(count)}</span>

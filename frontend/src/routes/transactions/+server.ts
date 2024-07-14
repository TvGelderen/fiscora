import { json } from '@sveltejs/kit';

export async function POST({ request }) {
    console.log(request);

    const formData = await request.formData();

    console.log(formData);

    return json({ success: true });
}
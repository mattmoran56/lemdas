export const getOptions = (method) => {
	const token = window.localStorage.getItem("token");

	const headers = new Headers();
	headers.append("Authorization", `Bearer ${token}`);
	headers.append("Content-Type", "application/json");

	return {
		method,
		headers,
	};
};

export const getOptionsWithBody = (method, body) => {
	const token = window.localStorage.getItem("token");

	const headers = new Headers();
	headers.append("Authorization", `Bearer ${token}`);

	return {
		method,
		headers,
		body,
	};
};

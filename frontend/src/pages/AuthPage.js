import React, { useEffect } from "react";
import { useSearchParams, useNavigate } from "react-router-dom";

const AuthPage = () => {
	const [searchParams] = useSearchParams();
	const navigate = useNavigate();

	const getAccessToken = async (code) => {
		try {
			const body = {
				code,
			};
			const requestOptions = {
				method: "POST",
				body: JSON.stringify(body),
			};
			const response = await fetch(
				`${process.env.REACT_APP_AUTH_API_URL}/token`,
				requestOptions,
			);
			if (!response.ok) {
				throw new Error("Failed to authorise user");
			}
			const data = await response.json();
			localStorage.setItem("token", data.access_token);
			navigate("/");
		} catch (err) {
			console.error(err);
		}
	};

	useEffect(() => {
		const code = searchParams.get("code");

		if (code !== "") {
			console.log(code);
			getAccessToken(code);
		}
	}, [searchParams]);

	return (
		<div className="w-screen h-screen bg-offwhite flex flex-col justify-center items-center">
			<h1 className="text-3xl mb-8 font-bold">Logging in through institution</h1>
			<p className="mb-12">We&apos;ll be one second...</p>
		</div>
	);
};

export default AuthPage;

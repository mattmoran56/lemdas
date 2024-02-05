import React, { useEffect, useState } from "react";
import { BuildingLibraryIcon } from "@heroicons/react/24/outline";
import { useNavigate } from "react-router-dom";

import Button from "../components/basic/Button";
import useAuth from "../hooks/useAuth";

const LoginPage = () => {
	const [usernameText, setUsernameText] = useState("");
	const [passwordText, setPasswordText] = useState("");

	const { username, error } = useAuth();

	const navigate = useNavigate();

	const handleLogin = () => {
		// TODO: handle login
	};

	const handleInstitutionLogin = () => {
		const domain = process.env.REACT_APP_AUTH_DOMAIN;
		const clientId = process.env.REACT_APP_AUTH_CLIENT_ID;
		// eslint-disable-next-line max-len
		window.location.href = `https://login.microsoftonline.com/${domain}/oauth2/v2.0/authorize?response_type=code&client_id=${clientId}&scope=User.Read`;
	};

	useEffect(() => {
		if (username && !error) {
			navigate("/");
		}
	}, [username, error, navigate]);

	return (
		<div className="w-screen h-screen bg-offwhite flex justify-center items-center">
			<div className="p-3 flex flex-col items-center w-80 max-w-full pb-8">
				<div className="text-indianred font-extrabold text-3xl mb-2">
					FYP
				</div>
				<h1 className="text-3xl mb-8 font-bold">Welcome back</h1>
				<input
					type="text"
					id="username"
					placeholder="Username"
					className="border-2 border-gray-400 rounded-3xl px-3 py-2 my-1 w-full bg-gray-300"
					value={usernameText}
					onChange={(event) => { return setUsernameText(event.target.value); }}
					disabled
				/>
				<input
					type="password"
					id="password"
					placeholder="Password"
					className="border-2 border-gray-400 rounded-3xl px-3 py-2 my-1 w-full bg-gray-300"
					value={passwordText}
					onChange={(event) => { return setPasswordText(event.target.value); }}
					disabled
				/>
				<Button
					className="border-gray-500 w-full my-3 bg-gray-500"
					onClick={handleLogin}
					disabled
				>
					Login
				</Button>
				<Button
					className="bg-transparent border-oxfordblue text-oxfordblue w-full mt-1"
					onClick={handleInstitutionLogin}
				>
					<BuildingLibraryIcon className="w-4 h-w mr-2" /> Login with Insitution
				</Button>
			</div>
		</div>
	);
};

export default LoginPage;

import React from "react";

const ErrorPage = () => {
	return (
		<div className="w-screen h-screen flex flex-col justify-center items-center pb-16">
			<div>
				<h1 className="text-7xl font-bold mb-8">404: <span className="text-indianred">Page not found</span></h1>
				<p className="text-xl">
					Looks like that page can&apos;t be found.
					Head <a href="/" className="underline">home</a> and try again!
				</p>
			</div>
		</div>
	);
};

export default ErrorPage;

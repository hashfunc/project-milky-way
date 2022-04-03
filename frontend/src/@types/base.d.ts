interface DefaultResponse<T, U = undefined> {
	code: string;
	message: string;
	meta?: U;
	data: T;
}

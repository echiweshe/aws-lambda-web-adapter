use http::{HeaderMap, Response};
use hyper::Body;

fn sanitize_headers<T>(response: &mut Response<T>) {
    let disallowed = [
        "connection",
        "keep-alive",
        "proxy-connection",
        "transfer-encoding",
        "upgrade",
    ];
    let headers = response.headers_mut();
    for name in disallowed.iter() {
        headers.remove(*name);
    }
}

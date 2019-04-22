import java.io.IOException;
import java.nio.charset.Charset;
import java.util.Map;
import java.util.TreeMap;
import okhttp3.MediaType;
import okhttp3.RequestBody;
import okhttp3.internal.Util;
import okio.BufferedSink;

/**
 * 为了使用签名这里重新封装请求体
 */
public class JsonRequestBody extends RequestBody {
    private final String mJSON;
    private final byte[] mBytes;
    private Charset charset = Util.UTF_8;

    /**
     * 如果是map对象那么请传TreeMap
     *
     * @param data
     */
    public JsonRequestBody(Object data) {
        if (data instanceof Map) {
            data = new TreeMap((Map) data);
        }
        this.mJSON = App.getGson().toJson(data);
        this.mBytes = mJSON.getBytes(charset);
    }

    @Override
    public long contentLength() throws IOException {
        return mBytes.length;
    }

    public String getBody() {
        return mJSON;
    }

    @Override
    public MediaType contentType() {
        return MediaType.parse("application/json; charset=utf-8");
    }

    @Override
    public void writeTo(BufferedSink sink) throws IOException {
        sink.write(mBytes);
    }
}
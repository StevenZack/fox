package io.gitee.stevenzack.foxui.FObject.Widget;

import android.os.Build;
import android.support.v4.view.ViewCompat;
import android.util.Log;
import android.view.ViewGroup;
import android.webkit.DownloadListener;
import android.webkit.WebResourceRequest;
import android.webkit.WebView;
import android.webkit.WebViewClient;

import org.json.JSONObject;
import org.json.JSONTokener;

import java.util.HashMap;
import java.util.Iterator;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FWebView extends FObject {
    public WebView v;

    public FWebView(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new WebView(parentController);
        v.setId(ViewCompat.generateViewId());
        view=v;
        v.getSettings().setDefaultTextEncodingName("UTF -8");
        parentController.onDestroyEvent.add(new Runnable() {
            @Override
            public void run() {
                if (v!=null) {
                    ViewGroup parent = (ViewGroup) v.getParent();
                    if (parent != null) {
                        parent.removeView(v);
                    }
                    v.stopLoading();
                    v.getSettings().setJavaScriptEnabled(false);
                    v.clearHistory();
                    v.clearView();
                    v.removeAllViews();
                    v.destroy();
                    v = null;
                }
                parentController.viewmap.remove(vid);
            }
        });
    }

    @Override
    public String getAttr(String attr) {
        return super.getAttr(attr);
    }

    @Override
    public String setAttr(String attr, final String value, final String value2) {
        String str = super.setAttr(attr, value, value2);
        if (str != null) {
            return str;
        }
        switch (attr) {
            case "Uri":
                v.loadUrl(value);
                break;
            case "Focusable":
                v.setFocusable(value.equals("true"));
                break;
            case "SupportZoom":
                v.getSettings().setSupportZoom(value.equals("true"));
                break;
            case "BuiltInZoomControls":
                v.getSettings().setBuiltInZoomControls(value.equals("true"));
                break;
            case "UseWideViewPort":
                v.getSettings().setUseWideViewPort(value.equals("true"));
                break;
            case "AllowContentAccess":
                v.getSettings().setAllowContentAccess(value.equals("true"));
                break;
            case "AllowFileAccess":
                v.getSettings().setAllowFileAccess(value.equals("true"));
                break;
            case "AllowFileAccessFromFileURLs":
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                    v.getSettings().setAllowFileAccessFromFileURLs(value.equals("true"));
                }
                break;
            case "AllowUniversalAccessFromFileURLs":
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                    v.getSettings().setAllowUniversalAccessFromFileURLs(value.equals("true"));
                }
                break;
            case "AppCacheEnabled":
                v.getSettings().setAppCacheEnabled(value.equals("true"));
                break;
            case "BlockNetworkImage":
                v.getSettings().setBlockNetworkImage(value.equals("true"));
                break;
            case "BlockNetworkLoads":
                v.getSettings().setBlockNetworkLoads(value.equals("true"));
                break;
            case "JavaScriptEnabled":
                v.getSettings().setJavaScriptEnabled(value.equals("true"));
                break;
            case "OffscreenPreRaster":
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
                    v.getSettings().setOffscreenPreRaster(value.equals("true"));
                }
                break;
            case "SaveFormData":
                v.getSettings().setSaveFormData(value.equals("true"));
                break;
            case "SupportMultipleWindows":
                v.getSettings().setSupportMultipleWindows(value.equals("true"));
                break;
            case "TextZoom":
                try {
                    v.getSettings().setTextZoom(Integer.parseInt(value));
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "UserAgentString":
                v.getSettings().setUserAgentString(value);
                break;
            // -------------------
            case "HandleUrl":
                MyWebViewClient wc=new MyWebViewClient();
                if (!value.equals("")) {
                    try {
                        wc.urlHandlersMap = new HashMap<>();
                        JSONObject object = (JSONObject) (new JSONTokener(value).nextValue());
                        Log.d(TAG, "setAttr: HandleUrl"+value);
                        for (Iterator<String> it = object.keys(); it.hasNext(); ) {
                            String key = it.next();
                            wc.urlHandlersMap.put(key, object.getString(key));
                        }
                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                }
                v.setWebViewClient(wc);
                break;
            case "DownloadListener":
                v.setDownloadListener(new DownloadListener() {
                    @Override
                    public void onDownloadStart(String url, String userAgent, String contentDisposition, String mimetype, long contentLength) {
                        Fox.triggerFunction(parentController,value, url,"","");
                    }
                });
                break;
            case "LoadHtmlData":
                String a = value.replace("#", "%23");
                v.loadData(a, "text/html; charset=UTF-8", null);
                break;
        }
        return "";
    }
    class MyWebViewClient extends WebViewClient {
        public HashMap<String,String> urlHandlersMap=null;
        @Override
        public boolean shouldOverrideUrlLoading(WebView view, WebResourceRequest request) {
            if (urlHandlersMap==null) {
                return super.shouldOverrideUrlLoading(view, request);
            }else {
                Log.d(TAG, "shouldOverrideUrlLoading: "+urlHandlersMap.size());
                if (android.os.Build.VERSION.SDK_INT >= android.os.Build.VERSION_CODES.LOLLIPOP) {
                    String url = request.getUrl().toString();
                    if (urlHandlersMap.containsKey(removeQuery(url))) {
                        String fnId = urlHandlersMap.get(removeQuery(url));
                        String s= Fox.triggerFunction(parentController,fnId, url,"","");
                        if(s!=null&&s.equals("true"))
                            return true;
                    }else if (urlHandlersMap.containsKey(url)) {
                        String fnId = urlHandlersMap.get(url);
                        String s= Fox.triggerFunction(parentController,fnId, url,"","");
                        if(s!=null&&s.equals("true"))
                            return true;
                    }
                }
            }
            return false;
        }

        @Override
        public boolean shouldOverrideUrlLoading(WebView view, String url) {
            if (urlHandlersMap==null) {
                return super.shouldOverrideUrlLoading(view, url);
            }else {
                Log.d(TAG, "shouldOverrideUrlLoading2: "+urlHandlersMap.size());
                if (android.os.Build.VERSION.SDK_INT >= android.os.Build.VERSION_CODES.LOLLIPOP) {
                    if (urlHandlersMap.containsKey(removeQuery(url))) {
                        String fnId = urlHandlersMap.get(removeQuery(url));
                        String s= Fox.triggerFunction(parentController,fnId, url,"","");
                        if(s!=null&&s.equals("true"))
                            return true;
                    }else if (urlHandlersMap.containsKey(url)) {
                        String fnId = urlHandlersMap.get(url);
                        String s= Fox.triggerFunction(parentController,fnId, url,"","");
                        if(s!=null&&s.equals("true"))
                            return true;
                    }
                }
            }
            return false;
        }
    }

    public static String removeQuery(String url) {
        int i = url.indexOf("?");
        if (i < 0) {
            return url;
        }
        return url.substring(0, i);
    }
}

package io.gitee.stevenzack.foxui;

import android.content.ContentUris;
import android.content.Context;
import android.content.Intent;
import android.database.Cursor;
import android.graphics.Bitmap;
import android.graphics.Canvas;
import android.graphics.drawable.BitmapDrawable;
import android.graphics.drawable.Drawable;
import android.net.Uri;
import android.os.Build;
import android.os.Bundle;
import android.os.Environment;
import android.os.Parcelable;
import android.provider.DocumentsContract;
import android.provider.MediaStore;
import android.support.annotation.NonNull;
import android.support.annotation.RequiresApi;
import android.support.v4.content.ContextCompat;
import android.support.v4.content.FileProvider;
import android.support.v4.graphics.drawable.DrawableCompat;
import android.support.v4.view.ViewCompat;
import android.util.Log;
import android.view.Menu;
import android.view.MenuItem;
import android.webkit.MimeTypeMap;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.BufferedInputStream;
import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.URISyntaxException;
import java.net.URL;
import java.net.URLConnection;
import java.util.ArrayList;
import java.util.List;


import static android.content.ContentValues.TAG;

public class Toolkit {
    public static void parseMenu(final FoxActivity uiController, Menu menu, JSONArray array) throws JSONException {
        for (int i=0;i<array.length();i++) {
            JSONObject object = array.getJSONObject(i);
            if (!object.has("MySubMenu")||object.isNull("MySubMenu")) {
                final MenuItem item = menu.add(0, ViewCompat.generateViewId(),i,object.getString("MyTitle"));
                if (!object.isNull("MyOnClick") && !object.getString("MyOnClick").equals("")) {
                    uiController.menuItemsOnClickMap.put(item, object.getString("MyOnClick"));
                }
                if (!object.isNull("MyIcon") && !object.getString("MyIcon").equals("")) {
                    String mIcon=object.getString("MyIcon");
                    file2Drawable(uiController, mIcon, new OnDrawableReadyListener() {
                        @Override
                        public void onDrawableReady(Drawable drawable) {
                            item.setIcon(drawable);
                        }
                    });
                }
                if (!object.isNull("MyShowAsAction") && !object.getString("MyShowAsAction").equals("")) {
                    item.setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
                }
                continue;
            }
            JSONArray subMenu=object.getJSONArray("MySubMenu");
            parseMenu(uiController,menu.addSubMenu(object.getString("MyTitle")),subMenu);
        }
    }
    public static String getPath(Context context, Uri uri) throws URISyntaxException {
        if ("content".equalsIgnoreCase(uri.getScheme())) {
            String[] projection = {"_data"};
            Cursor cursor = null;
            try {
                cursor = context.getContentResolver().query(uri, projection, null, null, null);
                assert cursor != null;
                int column_index = cursor.getColumnIndexOrThrow("_data");
                if (cursor.moveToFirst()) {
                    return cursor.getString(column_index);
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        } else if ("file".equalsIgnoreCase(uri.getScheme())) {
            return uri.getPath();
        }
        return null;
    }

    public static String getPathByUriBelowKitkat(Context context, Uri data) {
        String filename = null;
        if (data.getScheme().toString().compareTo("content") == 0) {
            Cursor cursor = context.getContentResolver().query(data, new String[]{"_data"}, null, null, null);
            if (cursor.moveToFirst()) {
                filename = cursor.getString(0);
            }
        } else if (data.getScheme().toString().compareTo("file") == 0) {// file:///开头的uri
            filename = data.toString();
            filename = data.toString().replace("file://", "");// 替换file://
            if (!filename.startsWith("/mnt")) {// 加上"/mnt"头
                filename += "/mnt";
            }
        }
        return filename;
    }

    // 专为Android4.4设计的从Uri获取文件绝对路径，以前的方法已不好使
    @RequiresApi(api = Build.VERSION_CODES.KITKAT)
    public static String getPathByUri4kitkat(final Context context, final Uri uri) {
        final boolean isKitKat = Build.VERSION.SDK_INT >= Build.VERSION_CODES.KITKAT;
        // DocumentProvider
        if (isKitKat && DocumentsContract.isDocumentUri(context, uri)) {
            if (isExternalStorageDocument(uri)) {// ExternalStorageProvider
                final String docId = DocumentsContract.getDocumentId(uri);
                final String[] split = docId.split(":");
                final String type = split[0];
                if ("primary".equalsIgnoreCase(type)) {
                    return Environment.getExternalStorageDirectory() + "/" + split[1];
                }
            } else if (isDownloadsDocument(uri)) {// DownloadsProvider
                final String id = DocumentsContract.getDocumentId(uri);
                final Uri contentUri = ContentUris.withAppendedId(Uri.parse("content://downloads/public_downloads"),
                        Long.valueOf(id));
                return getDataColumn(context, contentUri, null, null);
            } else if (isMediaDocument(uri)) {// MediaProvider
                final String docId = DocumentsContract.getDocumentId(uri);
                final String[] split = docId.split(":");
                final String type = split[0];
                Uri contentUri = null;
                if ("image".equals(type)) {
                    contentUri = MediaStore.Images.Media.EXTERNAL_CONTENT_URI;
                } else if ("video".equals(type)) {
                    contentUri = MediaStore.Video.Media.EXTERNAL_CONTENT_URI;
                } else if ("audio".equals(type)) {
                    contentUri = MediaStore.Audio.Media.EXTERNAL_CONTENT_URI;
                }
                final String selection = "_id=?";
                final String[] selectionArgs = new String[]{split[1]};
                return getDataColumn(context, contentUri, selection, selectionArgs);
            }
        } else if ("content".equalsIgnoreCase(uri.getScheme())) {// MediaStore
            // (and
            // general)
            return getDataColumn(context, uri, null, null);
        } else if ("file".equalsIgnoreCase(uri.getScheme())) {// File
            return uri.getPath();
        }
        return null;
    }

    public static String getDataColumn(Context context, Uri uri, String selection, String[] selectionArgs) {
        Cursor cursor = null;
        final String column = "_data";
        final String[] projection = {column};
        try {
            cursor = context.getContentResolver().query(uri, projection, selection, selectionArgs, null);
            if (cursor != null && cursor.moveToFirst()) {
                final int column_index = cursor.getColumnIndexOrThrow(column);
                return cursor.getString(column_index);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        if (cursor != null)
            cursor.close();
        return null;
    }

    public static boolean isExternalStorageDocument(Uri uri) {
        return "com.android.externalstorage.documents".equals(uri.getAuthority());
    }

    public static boolean isDownloadsDocument(Uri uri) {
        return "com.android.providers.downloads.documents".equals(uri.getAuthority());
    }

    public static boolean isMediaDocument(Uri uri) {
        return "com.android.providers.media.documents".equals(uri.getAuthority());
    }

    public static void saveDrawable(Drawable drawable, Bitmap.CompressFormat format, String path) {
        Bitmap bitmap = getBitmapFromDrawable(drawable);
        saveBitmap(bitmap, format, path);
    }

    public static Bitmap drawableToBitmap(Drawable drawable) {
        if (drawable == null)
            return null;
        return ((BitmapDrawable) drawable).getBitmap();
    }

    @NonNull
    static private Bitmap getBitmapFromDrawable(@NonNull Drawable drawable) {
        final Bitmap bmp = Bitmap.createBitmap(drawable.getIntrinsicWidth(), drawable.getIntrinsicHeight(), Bitmap.Config.ARGB_8888);
        final Canvas canvas = new Canvas(bmp);
        drawable.setBounds(0, 0, canvas.getWidth(), canvas.getHeight());
        drawable.draw(canvas);
        return bmp;
    }

    public static void saveBitmap(Bitmap bitmap, Bitmap.CompressFormat format, String path) {
        // 创建一个位于SD卡上的文件
        File file = new File(path);
        try {
            if (!file.exists()) {
                file.createNewFile();
            }
            FileOutputStream out = null;
            // 打开指定文件输出流
            out = new FileOutputStream(file);
            // 将位图输出到指定文件
            bitmap.compress(format, 100,
                    out);
            out.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    interface DownloadedListener {
        void onFailed(String error);

        void onSucceed(String fpath);
    }

    public static void openFile(Context context, String path) {
        Intent intent = new Intent(Intent.ACTION_VIEW);
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.N) {
            /* Android N 写法*/
            intent.setFlags(Intent.FLAG_GRANT_READ_URI_PERMISSION);
            Uri contentUri = FileProvider.getUriForFile(context, BuildConfig.APPLICATION_ID + ".fileProvider", new File(path));
            intent.setDataAndType(contentUri, getMimeType(path));
        } else {
            /* Android N之前的老版本写法*/
            intent.setDataAndType(Uri.fromFile(new File(path)), getMimeType(path));
            intent.setFlags(Intent.FLAG_ACTIVITY_NEW_TASK);
        }
        context.startActivity(intent);
    }

    public static String getMimeType(String url) {
        String type = null;
        String extension = MimeTypeMap.getFileExtensionFromUrl(url);
        if (extension != null) {
            type = MimeTypeMap.getSingleton().getMimeTypeFromExtension(extension);
        }
        return type;
    }

    public static List<String> parsePaths(Context context, Intent intent) {
        String action=intent.getAction();
        Bundle extras=intent.getExtras();
        List<String> paths = new ArrayList<>();
        if (Intent.ACTION_SEND.equals(action)) {
            Uri uri = (Uri) extras.getParcelable(Intent.EXTRA_STREAM);
            if (uri==null)return paths;
            String path= null;
            try {
                path = Toolkit.getPath(context,uri);
                paths.add(path);
                return paths;
            } catch (URISyntaxException e) {
                e.printStackTrace();
            }
        }else if (Intent.ACTION_VIEW.equals(action)){
            Uri uri=intent.getData();
            if (uri==null)return paths;
            String path= null;
            try {
                path = getPath(context,uri);
                paths.add(path);
                return paths;
            } catch (URISyntaxException e) {
                e.printStackTrace();
            }
        }else if (Intent.ACTION_SEND_MULTIPLE.equals(action)) {
            ArrayList<Parcelable> list = intent.getParcelableArrayListExtra(Intent.EXTRA_STREAM);
            if (list == null || list.size() == 0) {
                return paths;
            }
            for (Parcelable parcelable: intent.getParcelableArrayListExtra(Intent.EXTRA_STREAM)) {
                Uri uri = (Uri) parcelable;
                if (uri==null)return paths;
                String path= null;
                try {
                    path = getPath(context,uri);
                    paths.add(path);
                } catch (URISyntaxException e) {
                    e.printStackTrace();
                    continue;
                }
            }
            return paths;
        }
        return paths;
    }

    public static String handleIntent(FoxActivity foxActivity, Intent intent) throws JSONException {
        JSONObject jsonIntent = new JSONObject();
        jsonIntent.put("Action", intent.getAction());
        List<String> ps = Toolkit.parsePaths(foxActivity, intent);
        // paths
        JSONArray jsonArray = new JSONArray();
        for (int i = 0; i < ps.size(); i++) {
            jsonArray.put(ps.get(i));
        }
        jsonIntent.put("Paths", jsonArray);
        //extras
        Bundle bundle = intent.getExtras();
        JSONObject jsonObject = new JSONObject();
        if (bundle != null) {
            for (String key : bundle.keySet()) {
                jsonObject.put(key, String.valueOf(bundle.get(key)));
            }
        }
        jsonIntent.put("Extras", jsonObject);
        return jsonIntent.toString();
    }

    public static float dp2pixel(FoxActivity activity, float dps) {
        float pxs = dps *activity.getResources().getDisplayMetrics().density;
        return pxs;
    }
    public static float pixel2dp(FoxActivity activity,float pxs) {
        float dps = pxs/activity.getResources().getDisplayMetrics().density;
        return dps;
    }

    public interface OnDrawableReadyListener{
        void onDrawableReady(Drawable d);
    }
    public static void file2Drawable(final FoxActivity parentUIController, String value, final OnDrawableReadyListener listener) {
        if (value.startsWith("file://")) {
            String path=value.substring("file://".length());
            Drawable draw=Drawable.createFromPath(path);
            listener.onDrawableReady(draw);
        } else if (value.startsWith("assets://")) {
//            Drawable d = Drawable.createFromStream(getAssets().open("Cloths/btn_no.png"), null);
            String path = value.substring("assets://".length());
            try {
                Drawable drawable = Drawable.createFromStream(parentUIController.getAssets().open(path), null);
                listener.onDrawableReady(drawable);
            } catch (Exception e) {
                e.printStackTrace();
            }
        } else if (value.startsWith("drawable://")) {
//            listener.onDrawableReady(ContextCompat.getDrawable(parentUIController.activity, src));
        } else if (value.equals("RippleEffect")) {
            Log.d(TAG, "file2Drawable: set rippleEffect1");
            listener.onDrawableReady( parentUIController.getResources().getDrawable(R.drawable.ripple));
            Log.d(TAG, "file2Drawable: set rippleEffect");
        } else if (value.startsWith("fdrawable://")) {
            listener.onDrawableReady( parentUIController.drawableMap.get(value));
        }
    }
}

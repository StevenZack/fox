package io.gitee.stevenzack.foxui.FObject.Service;

import android.content.ClipData;
import android.content.ClipboardManager;
import android.content.Context;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FClipboard extends FObject {
    public static ClipboardManager v;

    public FClipboard(FoxActivity activity) {
        parentController = activity;
        v = (ClipboardManager) parentController.getSystemService(Context.CLIPBOARD_SERVICE);
    }

    @Override
    public String getAttr(String attr) {
        if (attr.startsWith("Item")) {
            try {
                int index = Integer.parseInt(attr.substring("Item".length()));
                return v.getPrimaryClip().getItemAt(index).getText().toString();
            } catch (Exception e) {
                e.printStackTrace();
                return e.toString();
            }
        }
        switch (attr) {
            case "ClipData":
                return v.getPrimaryClip().getItemAt(0).getText().toString();
            case "ClipCount":
                return String.valueOf(v.getPrimaryClip().getItemCount());
        }
        return null;
    }

    @Override
    public String setAttr(String attr, final String value,String value2) {
        if (value == null) {
            return "";
        }
        switch (attr) {
            case "OnChange":
                v.addPrimaryClipChangedListener(new ClipboardManager.OnPrimaryClipChangedListener() {
                    @Override
                    public void onPrimaryClipChanged() {
                        Fox.triggerFunction(parentController, value, "", "", "");
                    }
                });
                break;
            case "ClipData":
                v.setPrimaryClip(ClipData.newPlainText("new",value));
                break;
        }
        return "";
    }
}

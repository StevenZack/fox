package io.gitee.stevenzack.foxui.FObject.Widget;

import android.view.MenuItem;
import android.widget.PopupMenu;

import org.json.JSONArray;
import org.json.JSONTokener;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;
import io.gitee.stevenzack.foxui.Toolkit;

public class FPopupMenu extends FObject {
    public PopupMenu v;

    public FPopupMenu(FoxActivity activity, FObject anchorView) {
        parentController = activity;
        v = new PopupMenu(parentController, anchorView.view);
    }
    @Override
    public String getAttr(String attr) {
        return null;
    }

    @Override
    public String setAttr(String attr, String value,String value2) {
        if (value == null) {
            return "";
        }
        switch (attr) {
            case "Menus":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    Toolkit.parseMenu(parentController, v.getMenu(), array);
                    v.setOnMenuItemClickListener(new PopupMenu.OnMenuItemClickListener() {
                        @Override
                        public boolean onMenuItemClick(MenuItem item) {
                            if (parentController.menuItemsOnClickMap.containsKey(item)) {
                                Fox.triggerFunction(parentController,parentController.menuItemsOnClickMap.get(item), "","","");
                                return true;
                            }
                            return false;
                        }
                    });
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "Show":
                v.show();
                break;
            case "Dismiss":
                v.dismiss();
                break;
        }
        return "";
    }
}

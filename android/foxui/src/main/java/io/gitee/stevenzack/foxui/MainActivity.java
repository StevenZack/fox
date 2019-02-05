package io.gitee.stevenzack.foxui;

import android.support.v4.widget.DrawerLayout;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

import fox.Fox;
import fox.IActivity;

public class MainActivity extends FoxActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        mainFoxUI((DrawerLayout) findViewById(R.id.ctn));
    }
}
